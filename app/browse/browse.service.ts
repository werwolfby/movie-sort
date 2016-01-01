import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import "rxjs/add/operator/map";
import {Observable} from "rxjs"

export interface FileInfo {
    folder: string;
    name: string;
}

export interface FileLinkInfo extends FileInfo {
    links: FileInfo[];
    new_links: FileInfo[];
}

@Injectable()
export class BrowseService {
    constructor(private _http: Http) {
    }
    
    getFiles() : Observable<FileLinkInfo[]> {
        return this._http
            .get("api/files.json")
            .map(r => <FileLinkInfo[]> r.json());
    }
}
