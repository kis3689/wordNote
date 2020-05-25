import { Component, OnInit, OnDestroy, Inject } from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {WordService} from "../word.service";
import {Word} from "../word";
import {Subscription} from 'rxjs';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {InstantErrorStateMatcher} from "./instant-error-state.matcher";

@Component({
  selector: 'app-word-dialog',
  templateUrl: './word-dialog.component.html',
  styleUrls: ['./word-dialog.component.scss']
})
export class WordDialogComponent implements OnDestroy {
  controlGroup: FormGroup;
  errorStateMatcher = new InstantErrorStateMatcher();
  addSubscription: Subscription;
  updateSubscription: Subscription;
  deleteSubscription: Subscription;

  constructor(
    @Inject(MAT_DIALOG_DATA) public word: Word,
    public dialogRef: MatDialogRef<WordDialogComponent>,
    public service: WordService
  ) { 
    this.controlGroup = new FormGroup({
      name: new FormControl(word.Name, Validators.required),
      mean: new FormControl(word.Mean, Validators.required)
    });
  }

  save(): void {
    this.word.Name = this.formValue('name');
    this.word.Mean = this.formValue('mean');
    if (!this.word.Id) {
      this.addSubscription = this.service.add(this.word)
        .subscribe(this.dialogRef.close);
    } else {
      this.updateSubscription = this.service.update(this.word)
        .subscribe(this.dialogRef.close);
    }
  }

  delete(): void {
    this.deleteSubscription = this.service.delete(this.word.Id)
      .subscribe(this.dialogRef.close);
  }

  hasError(controlName: string, errorCode: string): boolean {
    return !this.controlGroup.valid && this.controlGroup.hasError(errorCode, [controlName]);
  }

  ngOnDestroy() {
    if (this.addSubscription) {
      this.addSubscription.unsubscribe();
    }
    if (this.updateSubscription) {
      this.updateSubscription.unsubscribe();
    }
    if (this.deleteSubscription) {
      this.deleteSubscription.unsubscribe();
    }
  }

  private formValue(controlName: string): any {
    return this.controlGroup.get(controlName).value;
  }

}
