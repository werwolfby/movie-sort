import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import "rxjs/add/operator/map";
import {Observable} from "rxjs"

export interface FileInfo {
    folder: string;
    name: string;
    links: string[];
}

@Injectable()
export class BrowseService {
    constructor(private _http: Http) {
    }
    
    getFiles() {
        return this._http
            .get("api/files.json")
            .map(r => <FileInfo[]> r.json());
    }
}
