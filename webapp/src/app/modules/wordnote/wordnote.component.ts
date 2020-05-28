import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';
import { WordnoteService } from '../wordnote.service';
import { Word } from '../word';

@Component({
  selector: 'app-wordnote',
  templateUrl: './wordnote.component.html',
  styleUrls: ['./wordnote.component.scss']
})
export class WordnoteComponent implements OnInit, OnDestroy {
  displayedColumns = ['Name', 'Mean'];
  dataSource: Word[] = [];
  getAllSubscription: Subscription;
  dialogSubscription: Subscription;

  constructor(public dialog: MatDialog, public service: WordnoteService) { }

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
