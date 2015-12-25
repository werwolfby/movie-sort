import {Component}   from "angular2/core";
import {JsonPipe}   from "angular2/common";
import {RouteConfig, Router, ROUTER_DIRECTIVES} from "angular2/router";
import {BrowseComponent} from "../browse/browse.component";
import {LogsComponent}   from "../logs/logs.component";

@Component({
    selector: 'ms-app',
    template: `
    <nav class="navbar navbar-inverse navbar-fixed-top">
        <div class="container">
            <div class="navbar-header">
                <a class="navbar-brand" href="#">Movie Sort</a>
            </div>
            <div id="navbar" class="collapse navbar-collapse">
                <ul class="nav navbar-nav">
                    <li><a [routerLink]="['Browse']">Browse</a></li>
                    <li><a [routerLink]="['Logs']">Logs</a></li>
                    <li><a>Execute</a></li>
                </ul>
            </div>
        </div>
    </nav>

    <div class="container-fluid">

        <div class="movie-sorter">
            <router-outlet></router-outlet>
        </div>

    </div>    
    `,
    directives: [ROUTER_DIRECTIVES],
    pipes: [JsonPipe],
})
@RouteConfig([
    {path: '/browse', name: 'Browse', component: BrowseComponent, useAsDefault: true},
    {path: '/logs',   name: 'Logs',   component: LogsComponent},
])
export class AppComponent {
    constructor(public currentRouter: Router) {
    }
}
