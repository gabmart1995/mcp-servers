import multer from 'multer';

const storage = multer.diskStorage({
    destination: (request, file, callback) => {
        callback(null, './uploads/images');
    },
    filename: (request, file, callback) => {
        const name = `project-${Date.now()}-${file.originalname}`;
        callback(null, name);
    }  
});

const upload = multer({ storage });

export default upload;