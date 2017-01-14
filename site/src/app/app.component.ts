import { Component } from '@angular/core';
import { PicturesService } from './pictures.service';
import { Picture } from './picture';
import {Location} from '@angular/common';

@Component({
  selector: 'app-root',
  providers: [PicturesService],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  pictureServices: PicturesService;
  picture: Picture;
  pictureURL: string;
  gaps = [
    { time: 0, name: "Minute" },
    { time: 600, name: "10 minutes" },
    { time: 3600, name: "1 heure" },
    { time: 21600, name: "6 heures" },
    { time: 86400, name: "1 jour" }
  ];
  gap = this.gaps[0];

  constructor(picturesService: PicturesService) {
    this.pictureServices = picturesService;
    this.picture = new Picture('', '', 0, 0, 0);
    this.getLastPicture();
  }

  getLastPicture() {
    this.pictureServices.getLastPicture()
      .subscribe(
      picture => {
        this.picture = picture
        this.pictureURL = "https://storage.googleapis.com/ants-photos/" + picture.FileName
      },
      err => {
        console.log(err);
      });
  }

  getPreviousPicture() {
    if (this.picture && this.gap) {
      this.pictureServices.getPreviousPicture(this.picture.DateTime - this.gap.time)
        .subscribe(
        picture => {
          if (picture) {
            this.picture = picture
            this.pictureURL = "https://storage.googleapis.com/ants-photos/" + picture.FileName
          }
        },
        err => {
          console.log(err);
        });
    }
  }

  getNextPicture() {
    if (this.picture) {
      this.pictureServices.getNextPicture(this.picture.DateTime + this.gap.time)
        .subscribe(
        picture => {
          if (picture) {
            this.picture = picture
            this.pictureURL = "https://storage.googleapis.com/ants-photos/" + picture.FileName
          }
        },
        err => {
          console.log(err);
        });
    }
  }
}
