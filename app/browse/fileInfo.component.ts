import {Component, Input, OnInit} from "angular2/core";
import {FileInfo} from "./browse.service";
import {Settings, FolderInfo, SettingsService} from "../root/settings.service";
import {Observable} from "rxjs/Observable";
import {BehaviorSubject} from "rxjs/subject/BehaviorSubject";
import "rxjs/add/operator/concat";

@Component({
    selector: 'file-info',
    template: '{{displayName}}'
})
export class FileInfoComponent implements OnInit {
    private _file : BehaviorSubject<FileInfo> = new BehaviorSubject<FileInfo>(null);
    private _absolutePath : BehaviorSubject<boolean> = new BehaviorSubject<boolean>(true);
    public displayName : string;
    
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
    
    private static _getPath(file: FileInfo, settings: Settings, absolutePath: boolean, inputFolders: FolderInfo[], outputFolders: FolderInfo[]) {
        if (!file) {
            return '';
        }
        var folderItems = absolutePath 
            ? (inputFolders.concat(outputFolders).filter(i => i.name == file.folder).pop() || {path:[]}).path
            : [file.folder];
        
        var items = [...folderItems, ...(file.path || []), file.name];
        return items.join(settings.pathSeparator);
    }
    
    ngOnInit() {
        var inputFolders = this._settingsService.inputFolders();
        var outputFolders = this._settingsService.outputFolders();
        
        this._settingsService.settings().combineLatest(this._file, this._absolutePath, inputFolders, outputFolders,
            (settings, file, absolutePath, inputFolders, outputFolders) => FileInfoComponent._getPath(file, settings, absolutePath, inputFolders, outputFolders))
            .subscribe(f => {
                this.displayName = f;
            });
    }
}
