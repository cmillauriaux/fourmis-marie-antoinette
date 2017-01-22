import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { Article } from './article';
import { Message } from './message';
import { Observable } from 'rxjs/Rx';

import { environment } from '../../environments/environment';

// Import RxJs required methods
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Injectable()
export class BlogService {
    constructor(private http: Http) { }

    isAuthorized(): Observable<Message> {
        return this.http.get(environment.serverURL + 'api/blog/isAuthorized')
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getArticles(): Observable<Article[]> {
        return this.http.get(environment.serverURL + 'api/blog/articles')
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getDetailsArticles(): Observable<Article[]> {
        return this.http.get(environment.serverURL + 'api/blog/articles/details')
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getArticle(articleID): Observable<Article> {
        return this.http.get(environment.serverURL + 'api/blog/article/' + articleID)
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    getArticleToEdit(articleID): Observable<Article> {
        return this.http.get(environment.serverURL + 'api/blog/article/' + articleID + '/edit')
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    deleteArticle(articleID): Observable<string> {
        return this.http.delete(environment.serverURL + 'api/blog/article/' + articleID)
            .map((res: Response) => res.toString())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    addArticle(article): Observable<Article> {
        let bodyString = JSON.stringify(article);
        let headers = new Headers({ 'Content-Type': 'application/json' });

        return this.http.post(environment.serverURL + 'api/blog/articles/add', bodyString)
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }

    updateArticle(article): Observable<Article> {
        let bodyString = JSON.stringify(article);
        let headers = new Headers({ 'Content-Type': 'application/json' });

        return this.http.put(environment.serverURL + 'api/blog/article/' + article.ID, bodyString)
            .map((res: Response) => res.json())
            .catch((error: any) => Observable.throw(error.json().error || 'Server error'));
    }
}