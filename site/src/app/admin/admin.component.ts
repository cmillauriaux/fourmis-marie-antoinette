import { Component } from '@angular/core';
import { BlogService } from '../blog/blog.service';
import { Message } from '../blog/message';
import { Article } from '../blog/article';

@Component({
    selector: 'admin',
    providers: [BlogService],
    templateUrl: './admin.component.html'
})
export class AdminComponent {
    blogService: BlogService;
    message: Message;
    articles: Article[];
    article: Article = new Article(0, "", "", "", "", false, "");

    constructor(blogService: BlogService) {
        this.blogService = blogService;
        blogService.isAuthorized()
            .subscribe(
            message => {
                this.message = message;
                if (this.message.IsAdmin) {
                    this.getArticles();
                }
            },
            err => {
                console.log(err);
            });
    }

    getArticles() {
        this.blogService.getArticles()
            .subscribe(
            articles => {
                this.articles = articles;
            }
            ,
            err => {
                console.log(err);
            });
    }

    saveNewArticle() {
        this.blogService.addArticle(this.article)
            .subscribe(
            article => {
                this.article = article;
                this.getArticles();
            }
            ,
            err => {
                console.log(err);
            });
    }

    saveArticle() {
        this.blogService.updateArticle(this.article)
            .subscribe(
            article => {
                this.article = article;
                this.getArticles();
            }
            ,
            err => {
                console.log(err);
            });
    }

    getArticle(articleID) {
        this.blogService.getArticle(articleID)
            .subscribe(
            article => {
                this.article = article;
            }
            ,
            err => {
                console.log(err);
            });
    }

    deleteArticle(articleID) {
        this.blogService.deleteArticle(articleID)
            .subscribe(
            message => {
                this.getArticles();
            }
            ,
            err => {
                console.log(err);
            });
    }
}