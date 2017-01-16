import { Component } from '@angular/core';
import { BlogService } from '../blog/blog.service';
import { Message } from '../blog/message';

@Component({
    selector: 'admin',
    providers: [BlogService],
    templateUrl: './admin.component.html'
})
export class AdminComponent {
    blogService: BlogService;
    message: Message;

    constructor(blogService: BlogService) {
        blogService.isAuthorized()
            .subscribe(
            message => {
                this.message = message
            },
            err => {
                console.log(err);
            });
    }
}