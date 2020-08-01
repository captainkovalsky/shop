import {Component, Input, OnInit} from '@angular/core';
import {Experience} from "../profile.service";

@Component({
  selector: 'app-experience',
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
            <p>{{exp?.role}}</p>

      </div>
      <div class="content extra">
      DESCRIPTION
      </div>
    </div>
  </article>
</div>
  `,
  styles: [`
    .extra {
      display: none;
    }
  `],
  animations: []
})
export class ExperienceComponent implements OnInit {
  @Input('exp') exp: Experience | null;

  constructor() {
    this.exp = null;
  }

  ngOnInit(): void {
  }
}
