import { Component } from '@angular/core';
import { Angulartics2GoogleAnalytics } from 'angulartics2';

@Component({
  selector: 'app-root',
  template: `
    <menu class="navbar"></menu>
    <router-outlet></router-outlet>
    <!--<footer></footer>-->
  `
})
export class AppComponent {
  constructor(angulartics2GoogleAnalytics: Angulartics2GoogleAnalytics) {}
}