import {Component, Input} from "angular2/core";
import {FileInfo, FileLinkInfo} from "./browse.service";
import {GuessitService, GuessitResult} from "./guessit.service";
import {TooltipDirective} from "../directives/tooltip.directive";

const showsFolder = "/shows";
const moviesFolder = "/movies";

@Component({
    selector: 'ms-guess-it',
    template: `
    <div [ngSwitch]="state">
        <template [ngSwitchWhen]="0"><a (click)="guessit()">Guess It</a></template>
        <template [ngSwitchWhen]="1">...loading...</template>
        <template [ngSwitchWhen]="2">
            <div *ngIf="!editLink">
                <button (click)="startEdit()" class="btn btn-primary btn-xs" type="submit">
                    <span class="glyphicon glyphicon-pencil" aria-hidden="true"></span>
                    <span>Edit</span>
                </button>
                <button (click)="cancel()" class="btn btn-warning btn-xs" type="submit">
                    <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                    <span>Cancel</span>
                </button>
                <span>{{newLink.folder}}/{{newLink.name}}</span>
                <button (click)="link()"  class="btn btn-success btn-xs" type="submit">
                    <span class="glyphicon glyphicon-ok"     aria-hidden="true"></span>
                    <span>Link</span>
                </button>
            </div>
            <div *ngIf="editLink">
                <div class="form-inline">
                    <div class="form-group">
                        <button (click)="cancelEdit()" class="btn btn-primary btn-xs" type="submit">
                            <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                            <span>Cancel</span>
                        </button>
                    </div>
                    <div class="form-group">
                        <div class="input-group input-group-sm">
                            <div class="input-group-btn">
                                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><span class="caret"></span> {{editRootFolder}}/</button>
                                <ul class="dropdown-menu">
                                    <li *ngFor="#f of rootFolders" [class.active]="editRootFolder == f"><a (click)="setRootFolder(f)">{{f}}/</a></li>
                                </ul>
                            </div>
                            <input type="text" class="form-control" [(ngModel)]="editLink.folder" placeholder="folder">
                            <div class="input-group-addon">/</div>
                            <input type="text" class="form-control" [(ngModel)]="editLink.name" placeholder="name">
                        </div>
                    </div>
                    <div class="form-group">
                        <button (click)="saveEdit()" class="btn btn-primary btn-xs" type="submit">
                            <span class="glyphicon glyphicon-ok" aria-hidden="true"></span>
                            <span>Save</span>
                        </button>
                    </div>
                </div>
            </div>
        </template>
    </div>
    `,
    providers: [GuessitService],
    directives: [TooltipDirective]
})
export class GuessItCompoenent {
    @Input() file: FileLinkInfo;
    private state: number = 0;
    private isEdit: boolean = false;
    private newLink: FileInfo;
    private editRootFolder: string;
    private editLink: FileInfo;
    private rootFolders: string[] = [moviesFolder, showsFolder];
    
    constructor(private _guessitService: GuessitService) {
    }
    
    guessit() {
        this.state = 1;
        this._guessitService
            .guess(this.file.name)
            .subscribe(data => this.onGuess(data));        
    }
    
    onGuess(data: GuessitResult) {
        this.state = 2;
        this.isEdit = false;
        var folder : string[] = [];
        switch (data.type) {
            case "episode":
                folder = ["/shows", data.title, "Season " + data.season];
                break;
            case "movie":
                folder = ["/movie"];
                break;
        }
        
        this.newLink = {folder: folder.join("/"), name: this.file.name};
    }
    
    link() {
    }
    
    startEdit() {
        this.editLink = $.extend(true, {}, this.newLink);
        for (var i = 0; i < this.rootFolders.length; i++) {
            var rootFolder = this.rootFolders[i];
            if (this.editLink.folder.startsWith(rootFolder)) {
                this.editRootFolder = rootFolder;
                this.editLink.folder = this.editLink.folder.slice(Math.min(rootFolder.length + 1, this.editLink.folder.length));
            }
        }
    }
    
    setRootFolder(rootFolder) {
        this.editRootFolder = rootFolder;
    }
    
    cancelEdit() {
        this.editLink = null;
    }
    
    saveEdit() {
        // remove starting slash
        if (this.editLink.folder.length > 0 && this.editLink.folder[0] != '/') {
            this.editLink.folder = '/' + this.editLink.folder;
        }
        this.newLink = {folder: this.editRootFolder + this.editLink.folder, name: this.editLink.name};
        this.editLink = null;
    }
    
    cancel() {
        this.state = 0;
    }
}