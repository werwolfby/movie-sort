import {Component}   from "angular2/core";
import {RouteConfig, Location, ROUTER_DIRECTIVES} from "angular2/router";
import {BrowseComponent} from "../browse/browse.component";
import {LogsComponent}   from "../logs/logs.component";
import {SettingsService} from "./settings.service";

@Component({
    selector: 'ms-app',
    template: `
    <nav class="navbar navbar-inverse navbar-fixed-top">
        <div class="container-fluid">
            <div class="navbar-header">
                <a class="navbar-brand" href="#">Movie Sort</a>
            </div>
            <div id="navbar" class="collapse navbar-collapse">
                <ul class="nav navbar-nav">
                    <li [class.active]="isLocationStartsWith('/browse')"><a [routerLink]="['Browse']">Browse</a></li>
                    <li [class.active]="isLocationStartsWith('/logs')"><a [routerLink]="['Logs']">Logs</a></li>
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
    providers: [SettingsService]
})
@RouteConfig([
    {path: '/browse', name: 'Browse', component: BrowseComponent, useAsDefault: true},
    {path: '/logs',   name: 'Logs',   component: LogsComponent},
])
export class AppComponent {
    constructor(private _location: Location) {
    }
    
    isLocationStartsWith(path: string) {
        return this._location.path().startsWith(path);
    }
}
