import { Component } from '@angular/core';
@Component({
  selector: 'app-root',
  template: `
    <menu class="navbar"></menu>
    <router-outlet></router-outlet>
    <!--<footer></footer>-->
  `
})
export class AppComponent {
}