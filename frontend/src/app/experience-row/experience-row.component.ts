import {Component, Input, OnInit} from '@angular/core';
import {Experience} from "../profile.service";

@Component({
  selector: 'app-experience-row',
  template: `

<div class="box">
  <article class="media">
    <div class="media-left">
      <figure class="image is-64x64">
        <img src="{{exp?.companyLogo}}" alt="Image">
      </figure>
    </div>
    <div class="media-content">
      <div class="content">
        <p>
          <strong>{{exp?.company}}</strong> - {{exp?.project}}
          <br>
        </p>
      </div>
      <nav class="level is-mobile">
        <div class="level-left">
          <a class="level-item" aria-label="reply">
          </a>
        </div>
      </nav>
    </div>
  </article>
</div>
  `,
  styles: []
})
export class ExperienceRowComponent implements OnInit {
  @Input('exp') exp: Experience | null;

  constructor() {
    this.exp = null;
  }

  ngOnInit(): void {
  }

}
