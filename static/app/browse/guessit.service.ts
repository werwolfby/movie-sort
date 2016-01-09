import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import "rxjs/add/operator/map";
import {Observable} from "rxjs"
import {FileInfo} from "./browse.service";

@Injectable()
export class GuessitService {
    constructor(private _http: Http) {
    }
    
    guess(fileInfo: FileInfo) : Observable<FileInfo> {
        return this._http
            .get(["/api/guess", fileInfo.folder, ...fileInfo.path, fileInfo.name].join("/"))
            .map(r => <FileInfo> r.json());
    }
};
