import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { Message } from './message';
import { Article } from './article';
import { Observable } from 'rxjs/Rx';

// Import RxJs required methods
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Injectable()
export class BlogService {
    constructor(private http: Http) { }

    isAuthorized(): Observable<Message> {
        return this.http.get('https://prototype-149014.appspot.com/api/pictures/last')
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getArticles(): Observable<Article[]> {
        return this.http.get('https://prototype-149014.appspot.com/api/pictures/last')
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getArticle(articleID): Observable<Article> {
        return this.http.get('https://prototype-149014.appspot.com/api/pictures/previous/' + timestamp)
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }
}