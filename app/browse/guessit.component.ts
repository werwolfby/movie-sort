import {Component, Input} from "angular2/core";
import {FileLinkInfo} from "./browse.service";
import {GuessitService, GuessitResult} from "./guessit.service";
import {TooltipDirective} from "../directives/tooltip.directive";

@Component({
    selector: 'ms-guess-it',
    template: `
    <div [ngSwitch]="state">
        <template [ngSwitchWhen]="0"><a (click)="guessit()">Guess It</a></template>
        <template [ngSwitchWhen]="1">...loading...</template>
        <template [ngSwitchWhen]="2">
            <button (click)="edit()" class="btn btn-primary btn-xs" type="submit">
                <span class="glyphicon glyphicon-pencil" aria-hidden="true"></span>
                <span>Edit</span>
            </button>
            <button (click)="cancel()" class="btn btn-warning btn-xs" type="submit">
                <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
                <span>Cancel</span>
            </button>
            <span>{{file.new_links[0].folder}}/{{file.new_links[0].name}}</span>
            <button (click)="link()" class="btn btn-success btn-xs" type="submit">
                <span class="glyphicon glyphicon-ok"     aria-hidden="true"></span>
                <span>Link</span>
            </button>
        </template>
    </div>
    `,
    providers: [GuessitService],
    directives: [TooltipDirective]
})
export class GuessItCompoenent {
    @Input() file: FileLinkInfo;
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
        
        if (!this.file.new_links) {
            this.file.new_links = [];
        }
        
        this.file.new_links.push({folder: folder.join("/"), name: this.file.name});
    }
    
    link() {
    }
    
    edit() {
    }
    
    cancel() {
        this.state = 0;
    }
}