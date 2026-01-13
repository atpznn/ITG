import { google } from 'googleapis';
const auth = new google.auth.GoogleAuth({
    keyFile: 'poc/credentials.json',
    scopes: ['https://www.googleapis.com/auth/spreadsheets'],
});

async function readSheetData() {
    try {
        const sheets = google.sheets({ version: 'v4', auth });
        const spreadsheetId = '1L5IxMPGrYATmMvBoShMZfZwB8f9wm8lqQbJ0FczML9Y';
        const response = await sheets.spreadsheets.values.get({
            spreadsheetId,
            range: 'Sheet1!A1:B10',
        });
        const rows = response.data.values
        if (rows == undefined || rows == null) return
        if (rows.length) {
            console.log('ข้อมูลใน Sheet:');
            rows.map((row) => {
                console.log(`${row[0]} - ${row[1]}`);
            });
        } else {
            console.log('ไม่พบข้อมูล');
        }
    } catch (err) {
        console.error('เกิดข้อผิดพลาด:', err);
    }
}

readSheetData();