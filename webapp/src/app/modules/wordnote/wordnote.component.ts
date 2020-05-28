import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { MatDialog } from '@angular/material/dialog';
import { WordnoteService } from './wordnote.service';
import { Word } from './word';
import { WordDialogComponent } from './word-dialog/word-dialog.component';

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
    this.openDialog(new Word());
  }

  openEditDialog(wd: Word) {
    this.openDialog(new Word(wd.Id, wd.Name, wd.Mean));
  }

  private openDialog(wd: Word): void {
    this.dialogSubscription = this.dialog
      .open(WordDialogComponent, {data: wd, minWidth: '30%'})
      .afterClosed().subscribe(() => this.loadStudentsList());
  }

  private loadStudentsList(): void {
    this.getAllSubscription = this.service.getAll()
      .subscribe(word => this.dataSource = word);
  }

  ngOnInit() {
    this.loadStudentsList();
  }

  ngOnDestroy() {
    this.getAllSubscription.unsubscribe();
    if (this.dialogSubscription) {
      this.dialogSubscription.unsubscribe();
    }
  }

}
