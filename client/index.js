#!/usr/bin/env node

const program = require('commander');
const { prompt } = require('inquirer');

const { query } = require('./query');
const { invoke } = require('./invoke');

program
  .version('0.0.1')
  .description('Medical Record management system');

const addSessionQuestions = [
    {
        type : 'input',
        name : 'nik',
        message : 'Masukkan NIK ...'
    },
    {
        type : 'input',
        name : 'nip',
        message : 'Masukkan NIP ...'
    },
    {
        type : 'input',
        name : 'treatment',
        message : 'Masukkan Treatment ...'
    },
    {
        type : 'input',
        name : 'medication',
        message : 'Masukkan Medication ...'
    },
];
program
  .command('addSession')
  .alias('as')
  .description('Add a session')
  .action(() => {
    prompt(addSessionQuestions).then(answers => {
      var args = [answers.nik, answers.nip, answers.treatment, answers.medication];
      invoke('addSession', 'user4', args);
    })
  });

const patientQuestions = [
    {
      type : 'input',
      name : 'nik',
      message : 'Masukkan NIK ...'
    },
    {
      type : 'input',
      name : 'name',
      message : 'Masukkan Nama ...'
    },
];

program
  .command('addPatient')
  .alias('ap')
  .description('Add a patient')
  .action(() => {
    prompt(patientQuestions).then(answers => {
      var args = [answers.nik, answers.name];
      invoke('addPatient', 'user4', args);})
  });

program
  .command('getRecord')
  .alias('g')
  .description('Get a record')
  .action(() => {
    prompt([{
      type : 'input',
      name : 'nik',
      message : 'Masukkan NIK ...'
    }]).then(answers =>
      query('queryByNIK', 'user4', answers.nik));
  });

program.parse(process.argv);
