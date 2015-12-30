import {Component, Input} from "angular2/core";
import {FileInfo} from "./browse.service";
import {GuessItCompoenent} from "./guessit.component";

@Component({
    selector: "ms-links-table",
    template: `
    <table class="table table-striped">
        <tr>
            <th>Source</th>
            <th>Dest</th>
        </tr>
        <tr *ngFor="#file of files">
            <td width="50%">{{file.folder}}/{{file.name}}</td>
            <td width="50%" *ngIf="file.links && file.links.length > 0"><div *ngFor="#link of file.links">{{link}}</div></td>
            <td width="50%" *ngIf="!file.links || file.links.length == 0"><ms-guess-it [file]="file"></ms-guess-it></td>
        </tr>
    </table>
    `,
    directives: [GuessItCompoenent]
})
export class LinksTableComponent {
    @Input() public files: FileInfo[];
}
