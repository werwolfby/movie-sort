import {Component, Input} from "angular2/core";
import {FileInfo} from "./browse.service";

@Component({
    selector: "ms-links-table",
    template: `
    <table class="table table-striped">
        <tr>
            <th>Source</th>
            <th>Dest</th>
        </tr>
        <tr *ngFor="#file of files">
            <td>{{file.name}}</td>
            <td><div *ngFor="#link of file.links">{{link}}</div></td>
        </tr>
    </table>
    `
})
export class LinksTableComponent {
    @Input() public files: FileInfo[]
}
