import {Component, Input} from "angular2/core";
import {FileInfo, FileLinkInfo, BrowseService} from "./browse.service";
import {GuessitService} from "./guessit.service";
import {FileInfoComponent} from "./fileInfo.component";
import {SettingsService} from "../root/settings.service";
import {TooltipDirective} from "../directives/tooltip.directive";

const showsFolder = "Shows";
const moviesFolder = "Movies";

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
                <file-info [file]="newLink"></file-info>
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
                                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><span class="caret"></span> {{editLink.rootFolder}}/</button>
                                <ul class="dropdown-menu">
                                    <li *ngFor="#f of rootFolders" [class.active]="editLink.folder == f"><a (click)="setRootFolder(f)">{{f}}/</a></li>
                                </ul>
                            </div>
                            <input type="text" class="form-control" [ngModel]="getPath(editLink)" (blur)="setPath(editLink, $event.target.value)" placeholder="folder">
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
        <template [ngSwitchWhen]="3">...linking...</template>
    </div>
    `,
    providers: [GuessitService],
    directives: [TooltipDirective, FileInfoComponent]
})
export class GuessItComponent {
    @Input() file: FileLinkInfo;
    private state: number = 0;
    private newLink: FileInfo;
    private editLink: FileInfo;
    private rootFolders: string[] = [moviesFolder, showsFolder];
    
    constructor(private _guessitService: GuessitService, private _browseService: BrowseService, private _settingsService: SettingsService) {
    }
    
    guessit() {
        this.state = 1;
        this._guessitService
            .guess(this.file)
            .subscribe(data => this.onGuess(data));
    }
    
    onGuess(data: FileInfo) {
        this.state = 2;
        this.newLink = data;
    }
    
    getPath(file: FileInfo) {
        if (!file.path || file.path.length == 0)
            return '';
        return file.path.join('/');
    }
    
    setPath(file: FileInfo, path: string) {
        if (path.length == 0) {
            return;
        }
        
        var folders = path.split('/').filter(p => p && p.length > 0);
        file.path = folders;
    }
    
    startEdit() {
        this.editLink = $.extend({}, this.newLink);
        /*this.editLink.rootFolder = this.rootFolders.filter(f => this.newLink.folder.startsWith(f)).pop() || null;
        if (this.editLink.rootFolder) {
            var length = Math.min(this.editLink.rootFolder.length + 1, this.editLink.relativeFolder.length);
            this.editLink.relativeFolder = this.editLink.relativeFolder.slice(length);
        }*/
    }
    
    setRootFolder(rootFolder) {
        this.editLink.folder = rootFolder;
    }
    
    cancelEdit() {
        this.editLink = null;
    }
    
    saveEdit() {
        this.newLink = this.editLink;
        this.editLink = null;
    }

    link() {
        this.state = 3;
        this._browseService.link(this.file, this.newLink)
            .subscribe(r => {
                if (!this.file.links) {
                    this.file.links = [];
                }
                this.file.links.push(...r.links);
                this.state = 0;
            });
    }
   
    cancel() {
        this.state = 0;
    }
}