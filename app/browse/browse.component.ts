import {Component, OnInit} from "angular2/core";
import {FileInfo, BrowseService} from "./browse.service";
import {LinksTableComponent} from "./links-table.component";

@Component({
    template: `
    <form>
        <div className="checkbox">
            <label>
                <input type="checkbox" [(ngModel)]="showWithoutLinksOnly"/>Show without links only
            </label>
        </div>
        <ms-links-table [files]="filteredFiles"></ms-links-table>
    </form>
    `,
    directives: [LinksTableComponent],
    providers: [BrowseService]
})
export class BrowseComponent implements OnInit {
    public filteredFiles: FileInfo[] = [];
    
    private _showWithoutLinksOnly = true;    
    private _allFiles: FileInfo[] = [];
    
    constructor(private _browseService: BrowseService) {        
    }
    
    get showWithoutLinksOnly() : boolean {
        return this._showWithoutLinksOnly;
    }
    
    set showWithoutLinksOnly(value: boolean) {
        this._showWithoutLinksOnly = value;
        this.filterFiles();
    }
    
    private filterFiles() {
        if (this.showWithoutLinksOnly) {
            this.filteredFiles = this._allFiles.filter(f => !f.links || f.links.length == 0);
        } else {
            this.filteredFiles = this._allFiles;
        }
    }
    
    ngOnInit() {
        this._browseService.getFiles().then(files => this._allFiles = files).then(() => this.filterFiles());
    }
}
