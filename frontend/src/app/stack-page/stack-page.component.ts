import {Component, Input, OnInit} from '@angular/core';
import {Experience, ProfileService, StackExperience} from "../profile.service";

export interface StackInfo {
  name: string;
  days: number;
  months: number;
}
@Component({
  selector: 'app-stack-page',
  template: `
          <div class="box" >
        <article class="media">
            <div class="table-container">
              <table class="table">
              <thead>
                <tr>
                  <th><abbr title="Technology">Technology</abbr></th>
                  <th>Days</th>
                  <th>Month</th>
                </tr>
              </thead>
              <tbody>
                <tr *ngFor="let info of stackInfo()">
                  <th>{{info.name}}
                  <td>{{info.days}}
                  <td>{{info.months}}
              </tbody>
                <!-- Your table content -->
              </table>
            </div>

        </article>
      </div>
  `,
  styles: []
})
export class StackPageComponent implements OnInit {
  @Input('exp') exp: Experience[];
  stack: StackExperience | null;

  constructor(private profile: ProfileService) {
    this.exp = [];
    this.stack = null;
  }

  ngOnInit(): void {
    this.stack = this.profile.GetStackExperience(this.exp);
  }

  stackInfo(): StackInfo[] {
    if (this.stack) {
      let stack = this.stack;
      return Object.keys(stack)
        .map(s => {
          // @ts-ignore
          return {name: s, days: stack[s].days, months: stack[s].months};
        })
        .sort(((a, b) => b.days - a.days));
    }

    return [];
  }

}
