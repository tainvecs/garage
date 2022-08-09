GARAGE_PY_API_SERVER_ROOT="$(dirname $(cd $(dirname $0) >/dev/null 2>&1; pwd -P;))"

python -m pip install --upgrade pip
python -m pip install -r $GARAGE_PY_API_SERVER_ROOT/requirements.txt

echo ">>> Please Create the .env file by referencing $GARAGE_PY_API_SERVER_ROOT/.env_template"
echo ">>> To source the .env file, you can "
echo ">>> run the command \"source $GARAGE_PY_API_SERVER_ROOT/.env\""
echo ">>> or simply use command line plugin 'autoenv'"
