import {Directive, ElementRef, OnInit, OnDestroy} from "angular2/core";

@Directive({
    selector: '[tooltip]'
})
export class TooltipDirective implements OnInit, OnDestroy {
    constructor(private _el: ElementRef) {
    }
    
    ngOnInit() {
        $(this._el.nativeElement).tooltip();
    }
    
    ngOnDestroy() {
        $(this._el.nativeElement).tooltip('destroy');
    }
}
