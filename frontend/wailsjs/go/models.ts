export namespace database {
	
	export class Entry {
	    name: string;
	    description: string;
	    url: string;
	    secret: string;
	
	    static createFrom(source: any = {}) {
	        return new Entry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	        this.url = source["url"];
	        this.secret = source["secret"];
	    }
	}

}

