import {Component} from "angular2/core";
import {LinksTableComponent} from "./links-table.component";

@Component({
    template: `
    <form>
        <div className="checkbox">
            <label>
                <input type="checkbox" [ngModel]="showWithoutLinksOnly" (ngModelChange)="showWithoutLinksOnly = $event"/>Show without links only
            </label>
        </div>
        <ms-links-table [withoutLinksOnly]="showWithoutLinksOnly"></ms-links-table>
    </form>
    `,
    directives: [LinksTableComponent]
})
export class BrowseComponent {
    showWithoutLinksOnly = true;    

    constructor() {
    }
}
