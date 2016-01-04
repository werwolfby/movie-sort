import {Component} from "angular2/core";
import {LinksTableComponent} from "./links-table.component";

@Component({
    template: `
    <form>
        <div class="container-fluid">
            <div class="pull-left">
                <div class="checkbox">
                    <label><input type="checkbox" [ngModel]="showWithoutLinksOnly" (ngModelChange)="showWithoutLinksOnly = $event">Show without links only</label>
                </div>
            </div>
            <div class="pull-right">
                <div class="checkbox">
                    <label><input type="checkbox" [ngModel]="showAbsolutePath" (ngModelChange)="showAbsolutePath = $event">Show absolute path</label>
                </div>
            </div>
        </div>
        <ms-links-table [withoutLinksOnly]="showWithoutLinksOnly" [absolutePath]="showAbsolutePath"></ms-links-table>
    </form>
    `,
    directives: [LinksTableComponent]
})
export class BrowseComponent {
    showWithoutLinksOnly = true;
    showAbsolutePath = false;

    constructor() {
    }
}
