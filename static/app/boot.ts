import {bootstrap}        from "angular2/platform/browser";
import {ROUTER_PROVIDERS} from "angular2/router";
import {HTTP_PROVIDERS}   from "angular2/http";
import {enableProdMode}   from "angular2/core";

import {AppComponent} from "./root/app.component";

import './../css/app.css';

enableProdMode();

bootstrap(AppComponent, [ROUTER_PROVIDERS, HTTP_PROVIDERS]);
