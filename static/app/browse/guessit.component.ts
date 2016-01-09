import {Component, Input, OnInit} from "angular2/core";
import {FileInfo, FileLinkInfo, BrowseService} from "./browse.service";
import {GuessitService} from "./guessit.service";
import {FileInfoComponent} from "./fileInfo.component";
import {SettingsService, Settings, FolderInfo} from "../root/settings.service";
import {TooltipDirective} from "../directives/tooltip.directive";
import {Observable} from "rxjs/Observable";
import {BehaviorSubject} from "rxjs/subject/BehaviorSubject";
import "rxjs/add/operator/combineLatest";

const showsFolder = "Shows";
const moviesFolder = "Movies";

@Component({
    selector: 'ms-guess-it',
    template: `
    <div [ngSwitch]="state">
        <template [ngSwitchWhen]="0"><a (click)="guessit()" class="btn btn-primary btn-xs"><span class="glyphicon glyphicon-link" aria-hidden="true"></span><span> Guess It</span></a></template>
        <template [ngSwitchWhen]="1">...loading...</template>
        <template [ngSwitchWhen]="2">
            <div *ngIf="!editLink">
                <button (click)="startEdit()" class="btn btn-primary btn-xs" type="submit">
                    <span class="glyphicon glyphicon-pencil" aria-hidden="true"></span>
                    <span> Edit</span>
                </button>
                <button (click)="cancel()" class="btn btn-warning btn-xs" type="submit">
                    <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                    <span> Cancel</span>
                </button>
                <file-info [file]="newLink" [absolutePath]="absolutePath"></file-info>
                <button (click)="link()"  class="btn btn-success btn-xs" type="submit">
                    <span class="glyphicon glyphicon-link"   aria-hidden="true"></span>
                    <span> Link</span>
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
                                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><span class="caret"></span> {{getFolderPathByName(editLink.folder)}}{{editSettings.pathSeparator}}</button>
                                <ul class="dropdown-menu">
                                    <li *ngFor="#f of rootFolders" [class.active]="editLink.folder == f.name"><a (click)="setRootFolder(f.name)">{{getFolderPath(f)}}{{editSettings.pathSeparator}}</a></li>
                                </ul>
                            </div>
                            <input type="text" class="form-control" [ngModel]="getPath(editLink)" (blur)="setPath(editLink, $event.target.value)" placeholder="folder">
                            <div class="input-group-addon">{{editSettings.pathSeparator}}</div>
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
export class GuessItComponent implements OnInit {
    @Input() file: FileLinkInfo;
    private state: number = 0;
    private newLink: FileInfo;
    private editLink: FileInfo;
    private editSettings: Settings;
    private rootFolders: FolderInfo[] = [];
    private _settings: Observable<Settings>;
    private _outputFolders: Observable<FolderInfo[]>;
    @Input()
    public absolutePath : boolean;
    
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
        return file.path.join(this.editSettings.pathSeparator);
    }
    
    setPath(file: FileInfo, path: string) {
        if (!path) {
            path ='';
        }
        
        var folders = path.split(this.editSettings.pathSeparator).filter(p => p && p.length > 0);
        file.path = folders;
    }
    
    startEdit() {
        this._outputFolders.combineLatest(this._settings, (outputFolders, settings) => ({outputFolders: outputFolders, settings: settings}))
            .subscribe(d => {
                this.rootFolders = d.outputFolders;
                this.editSettings = d.settings;
                this.editLink = $.extend({}, this.newLink);
            });
    }
    
    getFolderPathByName(folder: string) {
        if (this.absolutePath) {
            var folderInfo = this.rootFolders.filter(f => f.name == folder).pop();
            return folderInfo ? folderInfo.path.join(this.editSettings.pathSeparator) : folder;
        } else {
            return folder;
        }
    }
    
    getFolderPath(folder: FolderInfo) {
        if (this.absolutePath) {
            return folder.path.join(this.editSettings.pathSeparator);
        } else {
            return folder.name;
        }
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
    
    ngOnInit() {
        this._settings = this._settingsService.settings();
        this._outputFolders = this._settingsService.outputFolders();
    }
}