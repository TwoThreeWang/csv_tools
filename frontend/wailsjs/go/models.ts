export namespace main {
	
	export class Response {
	    data: string;
	    info: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.data = source["data"];
	        this.info = source["info"];
	    }
	}

}

