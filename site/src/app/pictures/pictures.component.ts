import { Component } from '@angular/core';
import { Location } from '@angular/common';
import { PicturesService } from './pictures.service';
import { Picture } from './picture';
import { Angulartics2 } from 'angulartics2';

@Component({
  selector: 'pictures',
  providers: [PicturesService],
  templateUrl: './pictures.component.html',
  styleUrls: ['./pictures.component.css']
})
export class PicturesComponent {
  angulartics2: Angulartics2;
  pictureServices: PicturesService;
  picture: Picture;
  pictureURL: string;
  gaps = [
    { time: 0, name: "1 Minute" },
    { time: 600, name: "10 minutes" },
    { time: 3600, name: "1 heure" },
    { time: 21600, name: "6 heures" },
    { time: 86400, name: "1 jour" }
  ];
  gap = this.gaps[0];

  constructor(picturesService: PicturesService, angulartics2: Angulartics2) {
    this.pictureServices = picturesService;
    this.angulartics2 = angulartics2;
    this.picture = new Picture('', '', 0, 0, 0);
    this.getLastPicture();
    angulartics2.eventTrack.next({ action: 'Last', properties: { category: 'Pictures' } });
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
    this.angulartics2.eventTrack.next({ action: 'Last', properties: { category: 'Pictures' } });
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
        this.angulartics2.eventTrack.next({ action: 'Previous', properties: { category: 'Pictures', label: this.gap.name } });
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
        this.angulartics2.eventTrack.next({ action: 'Next', properties: { category: 'Pictures', label: this.gap.name } });
    }
  }
}
