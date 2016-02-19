import {Directive, ElementRef, AfterViewInit, OnDestroy} from "angular2/core";

declare var $: any;

@Directive({
    selector: '[msTooltip]',
    inputs: ['title:msTooltip']
})
export class TooltipDirective implements AfterViewInit, OnDestroy {
    public get title(): string {
        return this._title;
    }
    public set title(value: string) {
        this._title = value;
        if (this._$el) {
            this._$el.attr('title', this._title);
            this._$el.tooltip('fixTitle');
        }
    }
    
    private _title: string;
    private _$el: any;
    
    constructor(private _el: ElementRef) {
    }
    
    ngAfterViewInit() {
        this._$el = $(this._el.nativeElement);
        this._$el.attr('data-toggle', 'tooltip');
        this._$el.attr('title', this._title);
        this._$el.tooltip();
    }
    
    ngOnDestroy() {
        $(this._el.nativeElement).tooltip('destroy');
    }
}
