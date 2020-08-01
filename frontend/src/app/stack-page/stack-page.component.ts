import {Component, Input, OnInit} from '@angular/core';
import {Experience, ProfileService, StackExperience} from "../profile.service";
import {add, formatDuration, intervalToDuration} from "date-fns";

export interface StackInfo {
  name: string;
  intervals: Duration[];
  sum: Duration;
}

@Component({
  selector: 'app-stack-page',
  template: `
          <div class="box">
        <article class="media">
            <div class="table-container">
              <table class="table is-striped is-fullwidth">
              <thead>
                <tr>
                  <th class="mobile"><abbr title="Technology">Technology</abbr></th>
                  <th>Duration</th>
                </tr>
              </thead>
              <tbody>
                <tr *ngFor="let info of stackInfo()">
                  <th>{{info.name}}
                  <td>{{format(info.sum)}}
              </tbody>
                <!-- Your table content -->
              </table>
            </div>
        </article>
      </div>
  `,
  styles: [
    `
    @media (max-width: 500px) {
      .mobile {
        width: 1%;
      }
    }
    `
  ]
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

  format(duration: Duration): string {
    return formatDuration(duration, {format: ['years', 'months']})
  }

  stackInfo(): StackInfo[] {
    if (this.stack) {
      let stack = this.stack;
      return Object.keys(stack)
        .map((s: string) => {
          return {
            name: s,
            intervals: stack[s],
            sum: stack[s].reduce((p, c) => {
              const start = new Date();
              const end = add(add(start, c), p);
              return intervalToDuration({start, end});
            }, {
              months: 1
            }),
          };
        })
        .sort((a, b) => {
          let first = (a.sum.days ?? 0) + (a.sum.years ?? 0) * 366 + (a.sum.months ?? 0) * 30;
          let second = (b.sum.days ?? 0) + (b.sum.years ?? 0) * 366 + (b.sum.months ?? 0) * 30;
          return second - first;
        });
    }

    return [];
  }
}
