
function testList {
  list="$1"

  curl localhost:8080/todo/lists/${list}/tasks -H "Authorization: Bearer $JWT"

  echo
  echo "----------------------"
  echo

  curl -X POST localhost:8080/todo/lists/${list}/tasks --data '{ "name" : "task 1", "description" : "this is task 1 for list'$list' " }' -H "Authorization: Bearer $JWT"

  curl -X POST localhost:8080/todo/lists/${list}/tasks --data '{ "name" : "task 2", "description" : "this is task 2" }' -H "Authorization: Bearer $JWT"

  curl -X POST localhost:8080/todo/lists/${list}/tasks --data '{ "name" : "task 3", "description" : "this is task 3" }' -H "Authorization: Bearer $JWT"

  echo
  echo "----------------------"
  echo

  curl localhost:8080/todo/lists/${list}/tasks -H "Authorization: Bearer $JWT"

}

lst=1
echo
echo "--------------------------------"
echo "Testing for list $lst"
testList $lst

lst=2
echo
echo "--------------------------------"
echo "Testing for list $lst"
testList $lst

