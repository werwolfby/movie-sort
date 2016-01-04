import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import "rxjs/add/operator/map";
import "rxjs/add/operator/publish";
import {Observable} from "rxjs/Observable"

export interface Settings {
    pathSeparator: string
}

export interface FolderInfo {
    name: string;
    pth: string[];
}

@Injectable()
export class SettingsService {
    private _settings: Observable<Settings>;
    private _inputFolders: Observable<FolderInfo[]>;
    private _outputFolders: Observable<FolderInfo[]>;
    
    constructor(private _http: Http) {
        this._settings = this._http
            .get("api/settings.json")
            .map(r => <Settings> r.json());
        this._inputFolders = this._http
            .get("api/settings/input-folders.json")
            .map(r => <FolderInfo[]> r.json());
        this._outputFolders = this._http
            .get("api/settings/output-folders.json")
            .map(r => <FolderInfo[]> r.json());
    }
    
    settings() : Observable<Settings> {
        return this._settings;
    }
    
    inputFolders() : Observable<FolderInfo[]> {
        return this._inputFolders;
    }
    
    outputFolders() : Observable<FolderInfo[]> {
        return this._outputFolders;
    }
}
