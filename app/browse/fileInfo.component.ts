import {Component, Input, OnInit} from "angular2/core";
import {FileInfo} from "./browse.service";
import {Settings, FolderInfo, SettingsService} from "../root/settings.service";
import {Observable} from "rxjs/Observable";
import {BehaviorSubject} from "rxjs/subject/BehaviorSubject";
import "rxjs/add/operator/concat";

interface FilePath {
    folderPath: string;
    filePath: string;
    isAbsolute: boolean;    
}

@Component({
    selector: 'file-info',
    template: '<span>{{displayName.folderPath}}</span>{{settings.pathSeparator}}<span>{{displayName.filePath}}</span>'
})
export class FileInfoComponent implements OnInit {
    private _file : BehaviorSubject<FileInfo> = new BehaviorSubject<FileInfo>(null);
    private _absolutePath : BehaviorSubject<boolean> = new BehaviorSubject<boolean>(true);
    public settings : Settings = {pathSeparator: '/'};
    public displayName : FilePath = {folderPath: '', filePath: '', isAbsolute: false};
    
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
            return {folderPath: '', filePath: '', isAbsolute: absolutePath};
        }
        var folderItems = absolutePath 
            ? (inputFolders.concat(outputFolders).filter(i => i.name == file.folder).pop() || {path:[]}).path
            : [file.folder];
        
        var items = [...(file.path || []), file.name];
        return {folderPath: folderItems.join(settings.pathSeparator), filePath: items.join(settings.pathSeparator), isAbsolute: absolutePath};
    }
    
    ngOnInit() {
        var inputFolders = this._settingsService.inputFolders();
        var outputFolders = this._settingsService.outputFolders();
        
        this._settingsService.settings().combineLatest(this._file, this._absolutePath, inputFolders, outputFolders,
            (settings, file, absolutePath, inputFolders, outputFolders) => ({
                settings: settings,
                displatName: FileInfoComponent._getPath(file, settings, absolutePath, inputFolders, outputFolders)
            }))
            .subscribe(f => {
                this.settings = f.settings;
                this.displayName = f.displatName;
            });
    }
}
