import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import "rxjs/add/operator/map";
import "rxjs/add/operator/mergeMap";
import {Observable} from "rxjs"

export interface FileInfo {
    folder: string;
    path: string[];
    name: string;
}

export interface FileLinkInfo extends FileInfo {
    links: FileInfo[];
}

@Injectable()
export class BrowseService {
    constructor(private _http: Http) {
    }
    
    getAll(folder?: string) : Observable<FileLinkInfo[]> {
        var path = "api/links"
        if (folder != null) {
            path = path + "/" + folder;            
        }
        return this._http
            .get(path)
            .map(r => <FileLinkInfo[]> r.json());
    }
    
    get(fileInfo: FileInfo) : Observable<FileInfo> {
        var path = "api/links/" + fileInfo.folder + "/" + fileInfo.path.join("/") + fileInfo.name;
        return this._http
            .get(path)
            .map(r => <FileInfo> r.json())
            .publish();
    }
    
    link(src, dest : FileInfo) : Observable<FileLinkInfo> {
        var path = "api/links/" + src.folder + "/" + src.path.join("/") + src.name;
        return this._http
            .put(path, JSON.stringify(dest))
            .map(r => {
                var result = <FileLinkInfo> r.json();
                result.links = [dest];
                return result;
            });
    }
}
