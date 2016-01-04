import {Component, Input, OnInit} from "angular2/core";
import {FileInfo, FileLinkInfo, BrowseService} from "./browse.service";
import {GuessItCompoenent} from "./guessit.component";
import {SettingsService, Settings} from "../root/settings.service";
import {Observable} from "rxjs/Observable";
import {BehaviorSubject} from "rxjs/subject/BehaviorSubject";
import "rxjs/add/operator/combineLatest";

@Component({
    selector: "ms-links-table",
    template: `
    <table class="table table-striped">
        <tr>
            <th width="50%">Source</th>
            <th width="50%">Dest</th>
        </tr>
        <tr *ngFor="#file of files">
            <td width="50%">{{fullPath(file)}}</td>
            <td width="50%" *ngIf="file.links && file.links.length > 0"><div *ngFor="#link of file.links">{{fullPath(link)}}</div></td>
            <td width="50%" *ngIf="!file.links || file.links.length == 0"><ms-guess-it [file]="file"></ms-guess-it></td>
        </tr>
    </table>
    `,
    directives: [GuessItCompoenent],
    providers: [BrowseService],
})
export class LinksTableComponent implements OnInit {
    private _settingsObservable: Observable<Settings>;
    private _allFilesObservable: Observable<FileLinkInfo[]>;
    private _withoutLinksOnlyObservable: BehaviorSubject<boolean> = new BehaviorSubject(false);
    
    private _settings : Settings;    
    public files : FileLinkInfo[] = [];
    
    @Input()
    public set withoutLinksOnly(value: boolean) {
        this._withoutLinksOnlyObservable.next(value);
    }
    
    constructor(private _settingsService: SettingsService, private _browseService : BrowseService) {
    }
    
    fullPath(file: FileInfo) {
        var items = [file.folder, ...(file.path || []), file.name];
        return items.join(this._settings ? this._settings.pathSeparator : "/");
    }
    
    ngOnInit() {
        this._settingsObservable = this._settingsService.settings();
        this._allFilesObservable = this._browseService.getAll();
        
        var settings = this._settingsObservable;
        var files = this._allFilesObservable
            .combineLatest(this._withoutLinksOnlyObservable, function(files, withoutLinksOnly){
                return files.filter(r => !withoutLinksOnly || ((r.links ? r.links.length : 0) == 0));
            });
        
        settings.combineLatest(files, function (settings, files) {
                return {settings: settings, files: files};
            })
            .subscribe(r => {
                console.log(JSON.stringify(r));
                this._settings = r.settings;
                this.files = r.files;
            });
    }
}
