import {Component, OnInit} from "angular2/core";
import {FileInfo, BrowseService} from "./browse.service";
import {LinksTableComponent} from "./links-table.component";
import {FilterPipe} from "../pipes/filter-func.pipe";

@Component({
    template: `
    <form>
        <div className="checkbox">
            <label>
                <input type="checkbox" [ngModel]="showWithoutLinksOnly" (ngModelChange)="showWithoutLinksOnly = $event; updateFilteredFiles()"/>Show without links only
            </label>
        </div>
        <ms-links-table [files]="filteredFiles"></ms-links-table>
    </form>
    `,
    pipes: [FilterPipe],
    directives: [LinksTableComponent],
    providers: [BrowseService]
})
export class BrowseComponent implements OnInit {
    showWithoutLinksOnly = true;    
    allFiles: FileInfo[] = [];
    filteredFiles: FileInfo[] = [];

    constructor(private _browseService: BrowseService) {        
    }
    
    private updateFilteredFiles() {
        if (this.showWithoutLinksOnly) {
            this.filteredFiles = this.allFiles.filter(f => !f.links || f.links.length == 0);
        } else {
            this.filteredFiles = this.allFiles;
        }
    }
    
    ngOnInit() {
        this._browseService.getFiles().then(files => this.allFiles = files).then(() => this.updateFilteredFiles());
    }
}
