import {Component, Input, OnInit} from "angular2/core";
import {FileInfo} from "./browse.service";
import {Settings, SettingsService} from "../root/settings.service";
import {Observable} from "rxjs/Observable";
import {BehaviorSubject} from "rxjs/subject/BehaviorSubject";

@Component({
    selector: 'file-info',
    template: '{{fullName}}'
})
export class FileInfoComponent implements OnInit {
    private _file : BehaviorSubject<FileInfo> = new BehaviorSubject<FileInfo>(null);
    public fullName : string;
    
    @Input()
    public set file(value: FileInfo) {
        this._file.next(value);
    }
    
    constructor(private _settingsService: SettingsService) {
    }
    
    private static _getFullPath(file: FileInfo, settings: Settings) {
        if (!file) {
            return '';
        }
        var items = [file.folder, ...(file.path || []), file.name];
        return items.join(settings.pathSeparator);
    }
    
    ngOnInit() {
        this._settingsService.settings().combineLatest(this._file, (settings, file) => FileInfoComponent._getFullPath(file, settings))
            .subscribe(f => {
                this.fullName = f;
            });
    }
}
