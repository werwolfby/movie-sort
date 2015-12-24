import {Component}   from "angular2/core";
import {RouteConfig, ROUTER_DIRECTIVES} from "angular2/router";
import {BrowseComponent} from "../browse/browse.component";

@Component({
    selector: 'ms-app',
    template: `<router-outlet></router-outlet>`,
    directives: [ROUTER_DIRECTIVES],
})
@RouteConfig([
    {path: '/browse', name: 'Browse', component: BrowseComponent, useAsDefault: true}
])
export class AppComponent {
}
