#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# Fetch LeetCode Questions

import argparse
import os
import json
import requests

LEETCODE_QUESTIONS_API="https://leetcode.com/api/problems/all/"
LEETCODE_GRAPHQL="https://leetcode.com/graphql"
DIFFICULTY=["None", "Easy", "Medium", "Hard"]

SOLUTIONS_FILE="solutions.md"

parser = argparse.ArgumentParser(description='Fetch LeetCode Questions.')
parser.add_argument('--id', metavar='id', help='question id', type=int)
parser.add_argument('-a', help='fetch all questions', action='store_true')
parser.add_argument('-f', help='finish one questions', action='store_true')
parser.add_argument('-n', help='start code', action='store_true')

args = parser.parse_args()

def get_questions():
    res = requests.get(LEETCODE_QUESTIONS_API)
    res = json.loads(res.content)
    questions = res['stat_status_pairs']
    def get_id(questions):
        return questions['stat']['question_id']
    questions.sort(key=get_id, reverse=False)
    return questions

def get_question_by_id(id):
    for q in get_questions():
        if q['stat']['question_id'] == id:
            return q

def fetch_all():
    f = open(SOLUTIONS_FILE, 'a')
    for q in get_questions():
        if not q['paid_only']:
            l = "| {0} | [{1}](https://leetcode.com/problems/{2}) | {3} | \n".format(q['stat']['question_id'], q['stat']['question__title'], q['stat']['question__title_slug'], DIFFICULTY[q['difficulty']['level']])
            f.write(l)
    f.close()

def finish_question(id):
    q = get_question_by_id(id)
    phase = "[{0}]".format(q['stat']['question__title'])
    line_number = 0
    f = open(SOLUTIONS_FILE,"r")
    for number, line in enumerate(f):
        if phase in line:
            line_number = number
            break
    f.close()

    if line_number <= 0:
        return

    f = open(SOLUTIONS_FILE,"r")
    lines = f.readlines()
    lines[line_number] = "| {0} | [{1}](https://leetcode.com/problems/{2}) | {3} | [Go](algorithms/{2}/{2}.go)\n".format(q['stat']['question_id'], q['stat']['question__title'], q['stat']['question__title_slug'], DIFFICULTY[q['difficulty']['level']])
    f.close()

    f2 = open(SOLUTIONS_FILE, "w")
    f2.writelines(lines)
    f2.close()

def new_solution(id):
    q = get_question_by_id(id)
    phase = "[{0}]".format(q['stat']['question__title'])
    data = {'operationName': 'questionData','variables': {'titleSlug': ''}, 'query': 'query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    exampleTestcases\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      paidOnly\n      hasVideoSolution\n      paidOnlyVideo\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    enableDebugger\n    envInfo\n    libraryUrl\n    adminUrl\n    __typename\n  }\n}\n'}
    data['variables']['titleSlug'] = q['stat']['question__title_slug']
    res = requests.post(LEETCODE_GRAPHQL, json= data)
    res = json.loads(res.content)
    q['question'] = res['data']['question']
    code = ""
    for c in q['question']['codeSnippets']:
        if c['lang'] == "Go":
            code = c['code']
            break

    file_content = "// https://leetcode.com/problems/{0}\n\npackage leetcode\n\n{1}\n".format(q['stat']['question__title_slug'], code)
    test_content = "package leetcode\n\nfunc Test{0}(t *testing.T) {1}".format(q['question']['title'].replace(" ", ""), "{\n\n}\n")
    filename = "algorithms/{0}/{0}.go".format(q['stat']['question__title_slug'])
    os.makedirs(os.path.dirname(filename), exist_ok=True)
    if not os.path.exists(filename):
        file = open(filename, "w")
        file.write(file_content)
        file.close()
    test_filename = "algorithms/{0}/{0}_test.go".format(q['stat']['question__title_slug'])
    if not os.path.exists(test_filename):
        file = open(test_filename, "w")
        file.write(test_content)
        file.close()

if args.a:
    fetch_all()

if args.f:
    finish_question(args.id)

if args.n:
    new_solution(args.id)

