import {Component, Input, OnInit} from "angular2/core";
import {FileInfo} from "./browse.service";
import {Settings, FolderInfo, SettingsService} from "../root/settings.service";
import {Observable} from "rxjs/Observable";
import {BehaviorSubject} from "rxjs/subject/BehaviorSubject";
import "rxjs/add/operator/concat";

interface FilePath {
    absoluteFolderPath: string;
    folderName: string;
    filePath: string;
    isAbsolute: boolean;    
}

@Component({
    selector: 'file-info',
    template: `<span class="folder-path text-primary" *ngIf="!displayName.isAbsolute" [title]="displayName.absoluteFolderPath">{{displayName.folderName}}{{settings.pathSeparator}}</span><span *ngIf="displayName.isAbsolute">{{displayName.absoluteFolderPath}}{{settings.pathSeparator}}</span><span>{{displayName.filePath}}</span>`
})
export class FileInfoComponent implements OnInit {
    private _file : BehaviorSubject<FileInfo> = new BehaviorSubject<FileInfo>(null);
    private _absolutePath : BehaviorSubject<boolean> = new BehaviorSubject<boolean>(true);
    public settings : Settings = {pathSeparator: '/'};
    public displayName : FilePath = {absoluteFolderPath: '', folderName: '', filePath: '', isAbsolute: false};
    
    @Input()
    public set file(value: FileInfo) {
        this._file.next(value);
    }
    
    @Input()
    public set absolutePath(value: boolean) {
        this._absolutePath.next(value);
    }
    
    constructor(private _settingsService: SettingsService) {
    }
    
    private static _getPath(file: FileInfo, settings: Settings, absolutePath: boolean, inputFolders: FolderInfo[], outputFolders: FolderInfo[]) : FilePath {
        if (!file) {
            return {absoluteFolderPath: '', folderName: '', filePath: '', isAbsolute: absolutePath};
        }
        var absoluteFolderPath = (inputFolders.concat(outputFolders).filter(i => i.name == file.folder).pop() || {path:[]}).path;
        
        var items = [...(file.path || []), file.name];
        return {absoluteFolderPath: absoluteFolderPath.join(settings.pathSeparator), folderName: file.folder, filePath: items.join(settings.pathSeparator), isAbsolute: absolutePath};
    }
    
    ngOnInit() {
        var inputFolders = this._settingsService.inputFolders();
        var outputFolders = this._settingsService.outputFolders();
        
        this._settingsService.settings().combineLatest(this._file, this._absolutePath, inputFolders, outputFolders,
            (settings, file, absolutePath, inputFolders, outputFolders) => ({
                settings: settings,
                displayName: FileInfoComponent._getPath(file, settings, absolutePath, inputFolders, outputFolders)
            }))
            .subscribe(f => {
                this.settings = f.settings;
                this.displayName = f.displayName;
            });
    }
}
