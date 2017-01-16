import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule } from '@angular/router';

import { AppComponent } from './app.component';
import { PicturesComponent } from './pictures/pictures.component';
import { ProjetComponent } from './projet/projet.component';
import { BlogComponent } from './blog/blog.component';
import { MenuComponent } from './menu/menu.component';
import { AdminComponent } from './admin/admin.component';
import { Angulartics2Module, Angulartics2GoogleAnalytics } from 'angulartics2';

@NgModule({
  declarations: [
    AppComponent,
    MenuComponent,
    PicturesComponent,
    ProjetComponent,
    BlogComponent,
    AdminComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    Angulartics2Module.forRoot([Angulartics2GoogleAnalytics]),
    RouterModule.forRoot([
      {
        path: 'pictures',
        component: PicturesComponent
      },
      {
        path: 'projet',
        component: ProjetComponent
      },
      {
        path: 'blog',
        component: BlogComponent
      },
            {
        path: 'admin',
        component: AdminComponent
      },
      {
        path: '',
        redirectTo: '/pictures',
        pathMatch: 'full'
      }
    ])
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
