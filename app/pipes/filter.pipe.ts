import {Pipe} from "angular2/core";

@Pipe({name: 'filter', pure: true})
export class FilterPipe {
    transform(value, [fn, args]) {
        return value.filter(v => fn(v, ...args));
    }
}
