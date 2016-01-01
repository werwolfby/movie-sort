import {Component, Input} from "angular2/core";
import {FileInfo} from "./browse.service";
import {GuessitService, GuessitResult} from "./guessit.service";
import {TooltipDirective} from "../directives/tooltip.directive";

@Component({
    selector: 'ms-guess-it',
    template: `
    <div [ngSwitch]="state">
        <template [ngSwitchWhen]="0"><a (click)="guessit()">Guess It</a></template>
        <template [ngSwitchWhen]="1">...loading...</template>
        <template [ngSwitchWhen]="2">
            <a (click)="edit()" tooltip data-toggle="tooltip" data-placement="bottom" title="Edit">
                <span class="glyphicon glyphicon-pencil" aria-hidden="true"></span>
            </a>
            <a (click)="cancel()">(cancel)</a>
            {{file.new_links[0]}}
            <a (click)="link()">(link)</a>               
        </template>
    </div>
    `,
    providers: [GuessitService],
    directives: [TooltipDirective]
})
export class GuessItCompoenent {
    @Input() file: FileInfo;
    state: number = 0;
    
    constructor(private _guessitService: GuessitService) {
    }
    
    guessit() {
        this.state = 1;
        this._guessitService
            .guess(this.file.name)
            .subscribe(data => this.onGuess(data));        
    }
    
    onGuess(data: GuessitResult) {
        this.state = 2;
        var folder : string[] = [];
        switch (data.type) {
            case "episode":
                folder = ["/shows", data.title, "Season " + data.season];
                break;
            case "movie":
                folder = ["/movie"];
                break;
        }
        folder.push(this.file.name);
        
        if (!this.file.new_links) {
            this.file.new_links = [];
        }
        
        this.file.new_links.push(folder.join("/"));
    }
    
    link() {
    }
    
    edit() {
    }
    
    cancel() {
        this.state = 0;
    }
}