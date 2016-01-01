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
            <span (click)="edit()"   tooltip data-toggle="tooltip" data-placement="bottom" title="Edit"   class="glyphicon glyphicon-pencil" aria-hidden="true"></span>
            <span (click)="cancel()" tooltip data-toggle="tooltip" data-placement="bottom" title="Cancel" class="glyphicon glyphicon-remove" aria-hidden="true"></span>
            <span>{{file.new_links[0]}}</span>
            <span (click)="link()"   tooltip data-toggle="tooltip" data-placement="bottom" title="Link"   class="glyphicon glyphicon-ok"     aria-hidden="true"></span>
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