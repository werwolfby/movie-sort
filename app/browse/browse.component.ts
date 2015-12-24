import {Component, OnInit} from "angular2/core";
import {FileInfo, BrowseService} from "./browse.service";
import {LinksTableComponent} from "./links-table.component";
import {FilterPipe} from "../pipes/filter.pipe";

@Component({
    template: `
    <form>
        <div className="checkbox">
            <label>
                <input type="checkbox" [(ngModel)]="showWithoutLinksOnly"/>Show without links only
            </label>
        </div>
        <ms-links-table [files]="allFiles | filter:filterFile:showWithoutLinksOnly"></ms-links-table>
    </form>
    `,
    pipes: [FilterPipe],
    directives: [LinksTableComponent],
    providers: [BrowseService]
})
export class BrowseComponent implements OnInit {
    showWithoutLinksOnly = true;    
    allFiles: FileInfo[] = [];
    
    constructor(private _browseService: BrowseService) {        
    }
    
    private filterFile(f: FileInfo, showWithoutLinksOnly: boolean) {        
        return !showWithoutLinksOnly || !f.links || f.links.length == 0;
    }
    
    ngOnInit() {
        this._browseService.getFiles().then(files => this.allFiles = files);
    }
}
