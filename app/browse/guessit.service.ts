import {Injectable} from "angular2/core";
import {Http} from "angular2/http";
import "rxjs/add/operator/map";
import {Observable} from "rxjs"

export interface GuessitResult {
    type: string;
    title: string;    
    season: number;
    episode: number;
};

@Injectable()
export class GuessitService {
    constructor(private _http: Http) {
    }
    
    guess(fileName: string) : Observable<GuessitResult> {
        return this._http
            .get("/api/guessit/guessit.json")
            .map(r => <GuessitResult> r.json());
    }
};
