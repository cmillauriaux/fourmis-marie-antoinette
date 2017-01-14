import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { Picture } from './picture';
import { Observable } from 'rxjs/Rx';

// Import RxJs required methods
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Injectable()
export class PicturesService {
    constructor(private http: Http) { }

    getLastPicture(): Observable<Picture> {
        return this.http.get('https://prototype-149014.appspot.com/api/pictures/last')
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getPreviousPicture(timestamp): Observable<Picture> {
        return this.http.get('https://prototype-149014.appspot.com/api/pictures/previous/' + timestamp)
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getNextPicture(timestamp): Observable<Picture> {
        return this.http.get('https://prototype-149014.appspot.com/api/pictures/next/' + timestamp)
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }
}