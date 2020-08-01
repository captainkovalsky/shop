import {Component, OnInit} from '@angular/core';
import {CV, ProfileService} from "./profile.service";

export enum Tab {
  EXP = 1,
  FAQ,
  CONTACT,
  STACK
}

@Component({
  selector: 'app-root',
  template: `
<section class="section">
  <div class="tabs is-medium">
    <ul (click)="SwitchTab($event)">
      <li [ngClass]="{'is-active': tab === tabs.EXP}"><a [attr.id]="tabs.EXP">Experience</a></li>
<!--      <li [ngClass]="{'is-active': tab === tabs.FAQ}"><a [attr.id]="tabs.FAQ">FAQs</a></li>-->
<!--      <li [ngClass]="{'is-active': tab === tabs.CONTACT}"><a [attr.id]="tabs.CONTACT">Contacts</a></li>-->
      <li [ngClass]="{'is-active': tab === tabs.STACK}"><a [attr.id]="tabs.STACK">Technologies</a></li>
    </ul>
  </div>

    <ng-template *ngIf="tab === tabs.EXP" ngFor let-exp [ngForOf]="cv?.experience">
          <app-experience [exp]="exp"></app-experience>
    </ng-template>

      <div class="box" *ngIf="tab === tabs.FAQ">
        <article class="media">
            FAQ
        </article>
      </div>

      <div class="box" *ngIf="tab === tabs.CONTACT">
        <article class="media">
            CONTACT
        </article>
      </div>

      <app-stack-page *ngIf="tab === tabs.STACK && cv && cv.experience" [exp]="cv.experience"></app-stack-page>

  <router-outlet></router-outlet>
</section>
  `,
  styles: []
})
export class AppComponent implements OnInit {
  title = 'CV';
  cv: CV | null = null;
  tab: Tab = Tab.EXP;
  tabs = Tab;

  constructor(private profileSrv: ProfileService) {

  }

  ngOnInit(): void {
    this.profileSrv.GetProfile().subscribe((profile: CV) => {
      this.cv = profile;
    })
  }

  SwitchTab($event: MouseEvent) {
    const {id = null} = $event.target as Element;

    if (!id) {
      return;
    }

    this.tab = +id;
  }
}
