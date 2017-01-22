import { Component } from '@angular/core';
import { BlogService } from './blog.service';
import { Article } from './article';

@Component({
  selector: 'blog',
  providers: [BlogService],
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.css']
})
export class BlogComponent {
  blogService: BlogService;
  articles: Article[];
  article: Article = new Article(0, "", "", "", "", false, "");

  constructor(blogService: BlogService) {
    this.blogService = blogService;
    this.getArticles();
  }

  getArticles() {
    this.blogService.getDetailsArticles()
      .subscribe(
      articles => {
        this.articles = articles;
      }
      ,
      err => {
        console.log(err);
      });
  }

  getArticle(articleID) {
    this.blogService.getArticleToEdit(articleID)
      .subscribe(
      article => {
        this.article = article;
      }
      ,
      err => {
        console.log(err);
      });
  }

}