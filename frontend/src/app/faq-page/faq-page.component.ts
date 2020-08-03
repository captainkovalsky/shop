import {Component, Input, OnInit} from '@angular/core';
import {Faq} from "../profile.service";


@Component({
  selector: 'app-faq-page',
  template: `
    <ng-template ngFor let-item [ngForOf]="faq" let-i="index">
        <article class="media is-size-5-desktop is-size-4-mobile">
        <div>
             <p class="has-text-weight-bold mb-4">{{item.question}}</p>
            <p [innerHTML]="item.answer"></p>
        </div>
        </article>
    </ng-template>

  `,
  styles: []
})
export class FaqPageComponent implements OnInit {
  @Input() faq: Faq[] = [];

  constructor() {
    this.faq = [];
  }

  ngOnInit(): void {
  }

}
