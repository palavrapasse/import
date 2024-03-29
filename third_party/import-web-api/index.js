import http from 'http';
import { Form } from 'multiparty';
import { exec } from 'child_process';
import * as dotenv from 'dotenv'
import { statSync } from 'fs';

dotenv.config();

const host = process.env.server_host;
const port = process.env.server_port;
const leaksDbFilePath = process.env.leaksdb_fp;
const subscribeNotifyUrl = process.env.subscribe_notify_url;

const corsHeaders = {
    'Access-Control-Allow-Origin': '*',
    'Access-Control-Allow-Methods': 'OPTIONS, POST',
};

const leakFormSchema = {
    context: new String(),
    leakFile: new String(),
    shareDate: new Date(),
    leakers: new String(),
    platforms: new String(),
};

function triggerImportLeak(leakForm, onFinish) {
    const cmd = `./import --database-path="${leaksDbFilePath}" --leak-path="${leakForm.leakFile}" --context="${leakForm.context}" --platforms="${leakForm.platforms}" --share-date="${leakForm.shareDate}" --leakers="${leakForm.leakers}" --notify-url="${subscribeNotifyUrl}" --skip=true`;
    console.info(cmd);

    exec(cmd, function (error, stdout, stderr) {
        console.log(`error!: ${error}`);
        console.log(`stdout!: ${stdout}`);
        console.log(`stderr!: ${stderr}`);

        onFinish();
    });
}

function leaksDbFileSizeInBytes() {
    return statSync(leaksDbFilePath).size;
}

const server = http.createServer(function (req, res) {
    if (req.method === 'POST') {
        const reqForm = new Form();

        reqForm.parse(req, function (err, fields, files) {
            if (!err) {
                const leakForm = Object.assign({}, leakFormSchema);

                leakForm.context = `${fields.context[0]}`;

                try {
                    leakForm.shareDate = new Date(Number.parseInt(`${fields.shareDateMS[0]}`)).toISOString().split('T')[0];
                    leakForm.platforms = `${fields.platforms[0]}`.split(',').filter((x) => x.length > 0);
                    leakForm.leakers = `${fields.leakers[0]}`.split(',').filter((x) => x.length > 0);
                    leakForm.leakFile = files.leakFile[0].path;

                    console.info(`triggering leak import with the following data: ${JSON.stringify(leakForm)}`);

                    const leaksDbFileSizeBeforeImport = leaksDbFileSizeInBytes();

                    triggerImportLeak(
                        leakForm,
                        function () {
                            const leaksDbFileSizeAfterImport = leaksDbFileSizeInBytes();
                            const sizeDiff = leaksDbFileSizeAfterImport - leaksDbFileSizeBeforeImport;

                            console.log(`size diff: ${sizeDiff} = ${leaksDbFileSizeAfterImport} - ${leaksDbFileSizeBeforeImport}`);

                            res.writeHead(200, { 'Content-Type': 'text/plain', ...corsHeaders }).end(`${sizeDiff}`);
                        }
                    );
                } catch (err) {
                    const error = `something went wrong!error: ${err}`;
                    console.error(error);

                    res.writeHead(500, { 'Content-Type': 'text/plain', ...corsHeaders }).end(error);
                }
            } else {
                console.error(err);

                res.writeHead(500, { 'Content-Type': 'text/plain', ...corsHeaders }).end(err);
            }
        });
    } else {
        res.writeHead(405, { 'Content-Type': 'text/plain' });
        res.end('Method Not Allowed\n');
    }
});

server.listen(port, host);
console.log(`Server running on port ${port}`);