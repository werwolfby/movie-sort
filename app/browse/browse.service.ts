import {Injectable} from "angular2/core";

let mockFiles : FileInfo[] = [
    {name: 'D:\\Video\\Complete\\Arrow.S01E08.rus.LostFilm.TV.avi', links: ['D:\\Video\\Serials\\Season 1\\Arrow.S01E08.rus.LostFilm.TV.avi']},
    {name: 'D:\\Video\\Complete\\Arrow.S01E09.rus.LostFilm.TV.avi', links: []},
    {name: 'D:\\Video\\Complete\\Пианистка DVDRip.avi', links: ['D:\\Video\\Films\\Пианистка DVDRip.avi']},
]

export interface FileInfo {
    name: string,
    links: string[]
}

@Injectable()
export class BrowseService {
    getFiles() {
        return Promise.resolve(mockFiles);
    }    
}
