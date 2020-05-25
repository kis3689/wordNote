import { Component, OnInit, OnDestroy } from '@angular/core';
import { Word } from './word'
import { Subscription } from 'rxjs'

import { MatDialog } from "@angular/material/dialog";
import { WordService } from './word.service'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit, OnDestroy {
  displayedColumns = ['Name', 'Mean'];
  dataSource: Word[] = [];
  getAllSubscription: Subscription;
  dialogSubscription: Subscription;

  constructor(public dialog: MatDialog, public service: WordService) {

  }

  openNewDialog() {
    alert('new')
  }

  openEditDialog(wd: Word) {
    alert('edit')
  }

  private loadStudentsList(): void {
    this.getAllSubscription = this.service.getAll()
      .subscribe(word => this.dataSource = word);
  }

  ngOnInit() {
    this.loadStudentsList();
  }

  ngOnDestroy() {

  }
}
